package gui

import (
	"database/sql"
	"fmt"
	"log"
	db_errors "remote_db/errors"
	"remote_db/usecase"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	_ "github.com/lib/pq"
)

func connectToDatabase(username, password string) (*sql.DB, error) {
	connStr := fmt.Sprintf("user=%s password=%s dbname=online_school host=localhost port=5432 sslmode=disable", username, password)
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil, fmt.Errorf("error connecting to the database: %v", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error pinging the database: %v", err)
	}

	return db, nil
}

func CreateMainWindow() fyne.Window {
	myApp := app.New()
	win := myApp.NewWindow("Online School")
	win.Resize(fyne.Size{Width: 650, Height: 500})

	showLoginForm(win)
	win.ShowAndRun()

	return win
}

func showLoginForm(win fyne.Window) {
	usernameEntry := widget.NewEntry()
	usernameEntry.SetPlaceHolder("Enter your username")

	passwordEntry := widget.NewPasswordEntry()
	passwordEntry.SetPlaceHolder("Enter your password")

	loginButton := widget.NewButton("Login", func() {
		db, err := connectToDatabase(usernameEntry.Text, passwordEntry.Text)
		if err != nil {
			dialog.ShowError(err, win)
			log.Println(err)
		} else {
			showMainMenu(win, db)
		}
	})

	content := container.NewVBox(
		usernameEntry,
		passwordEntry,
		loginButton,
	)
	win.SetContent(content)
}

func showMainMenu(win fyne.Window, db *sql.DB) {
	createTableButton := widget.NewButton("Create Table", func() {
		err := usecase.CreateStudentsTable(db)
		if err != nil {
			dialog.ShowError(err, win)
		} else {
			log.Println("Table created successfully")
		}
	})

	addTableButton := widget.NewButton("Add Student", func() {
		addRecordForm(win, db)
	})

	dropTableButton := widget.NewButton("Drop Table", func() {
		err := usecase.DropStudentsTable(db)
		if err != nil {
			dialog.ShowError(err, win)
			log.Println(err)
		} else {
			log.Println("Table was deleted successsfully")
		}
	})

	cleanTableButton := widget.NewButton("Clean Table", func() {
		err := usecase.TruncateStudentsTable(db)
		if err != nil {
			dialog.ShowError(err, win)
			log.Println(err)
		} else {
			log.Println("Table was cleaned")
		}
	})

	getRowsButton := widget.NewButton("Get students by course", func() { showStudentsByCourseForm(win, db) })
	deleteRowButton := widget.NewButton("Delete row by name", func() { deleteRowByNameForm(win, db) })

	upateEmailButton := widget.NewButton("Update email", func() { updateStudentEmailForm(win, db) })

	content := container.NewVBox(createTableButton, addTableButton, dropTableButton, cleanTableButton,
		deleteRowButton, upateEmailButton, getRowsButton)
	win.SetContent(content)
}

func addRecordForm(win fyne.Window, db *sql.DB) {
	name := widget.NewEntry()
	name.SetPlaceHolder("Enter name")
	email := widget.NewEntry()
	email.SetPlaceHolder("Enter email")
	course := widget.NewEntry()
	course.SetPlaceHolder("Enter course")

	backButton := widget.NewButton("Back", func() {
		showMainMenu(win, db)
	})

	addButton := widget.NewButton("Add new record", func() {
		if name.Text == "" || email.Text == "" {
			dialog.ShowError(db_errors.ErrEmptyValue, win)
			return
		}

		err := usecase.AddStudent(db, name.Text, email.Text, course.Text)
		if err != nil {
			log.Println("Error adding student:", err)
			dialog.ShowError(err, win)
		} else {
			log.Println("Student added successfully")
		}
	})

	form := container.NewVBox(name, email, course, addButton, backButton)

	win.SetContent(form)
}

func deleteRowByNameForm(win fyne.Window, db *sql.DB) {
	name := widget.NewEntry()
	name.SetPlaceHolder("Entry delete name")

	backButton := widget.NewButton("Back", func() {
		showMainMenu(win, db)
	})

	deleteRowButton := widget.NewButton("Delete row", func() {
		if name.Text == "" {
			dialog.ShowError(db_errors.ErrEmptyValue, win)
			return
		}
		err := usecase.DeleteStudentByName(db, name.Text)
		if err != nil {
			dialog.ShowError(err, win)
			log.Println(err)
		} else {
			log.Print("Recod was deleted successfully")
		}
	})

	form := container.NewVBox(name, backButton, deleteRowButton)
	win.SetContent(form)
}

func updateStudentEmailForm(win fyne.Window, db *sql.DB) {
	name := widget.NewEntry()
	name.SetPlaceHolder("Enter student name")
	email := widget.NewEntry()
	email.SetPlaceHolder("Enter new student email")

	backButton := widget.NewButton("Back", func() { showMainMenu(win, db) })

	updateButton := widget.NewButton("Update email", func() {
		if name.Text == "" || email.Text == "" {
			dialog.ShowError(db_errors.ErrEmptyValue, win)
			return
		}
		err := usecase.UpdateStudentEmail(db, name.Text, email.Text)
		if err != nil {
			log.Println(err)
			dialog.ShowError(err, win)
		} else {
			log.Println("Email was updated")
		}
	})

	form := container.NewVBox(name, email, backButton, updateButton)
	win.SetContent(form)
}

func showStudentsByCourseForm(win fyne.Window, db *sql.DB) {
	courseEntry := widget.NewEntry()
	courseEntry.SetPlaceHolder("Enter course name")

	studentList := widget.NewList(
		func() int { return 0 },
		func() fyne.CanvasObject { return widget.NewLabel("") },
		func(i widget.ListItemID, obj fyne.CanvasObject) {
			obj.(*widget.Label).SetText("")
		},
	)

	backButton := widget.NewButton("Back", func() {
		showMainMenu(win, db)
	})

	searchButton := widget.NewButton("Search", func() {
		course := courseEntry.Text
		if course == "" {
			dialog.ShowError(db_errors.ErrEmptyValue, win)
		}

		students, err := usecase.GetStudentsByCourse(db, course)
		if err != nil {
			dialog.ShowError(err, win)
			log.Println(err)
		}
		log.Println("Students found:", len(students))
		studentList.Length = func() int { return len(students) }
		studentList.UpdateItem = func(i widget.ListItemID, obj fyne.CanvasObject) {
			obj.(*widget.Label).SetText(fmt.Sprintf("%d. %s - %s", students[i].ID, students[i].Name, students[i].Email))
		}

		studentList.Refresh()
	})

	content := container.NewVBox(
		courseEntry,
		backButton,
		searchButton,
		studentList,
	)

	win.SetContent(content)
}
