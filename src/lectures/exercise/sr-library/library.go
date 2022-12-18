//--Summary:
//  Create a program to manage lending of library books.
//
//--Requirements:
//* The library must have books and members, and must include:
//  - Which books have been checked out
//  - What time the books were checked out
//  - What time the books were returned
//* Perform the following:
//  - Add at least 4 books and at least 3 members to the library
//  - Check out a book
//  - Check in a book
//  - Print out initial library information, and after each change
//* There must only ever be one copy of the library in memory at any time
//
//--Notes:
//* Use the `time` package from the standard library for check in/out times
//* Liberal use of type aliases, structs, and maps will help organize this project

package main

import (
	"fmt"
	"time"
)

const NA = "NA"

type CopiesRecord struct {
	total     int
	lent      int
	remaining int
}

type HistoryRecord struct {
	bookName    string
	memberName  string
	description string
	timeStamp   time.Time
}

type Library struct {
	name           string
	members        map[string]map[string]int
	books          map[string]CopiesRecord
	historyRecords []HistoryRecord
}

func printStatus(library *Library) {
	fmt.Println("=======================================")
	fmt.Println("Welcome to ", library.name, "Library!")
	fmt.Println("The current status of the library is:")
	fmt.Println("members:", library.members)
	fmt.Println("books:", library.books)
	fmt.Println("historyRecords:")
	for _, history := range library.historyRecords {
		fmt.Println(
			"Member:", history.memberName,
			"Book:", history.bookName,
			"Description:", history.description,
			"at time:", history.timeStamp)
	}
	fmt.Println("=======================================")
	fmt.Println()

}

func addMember(library *Library, memberName string) {
	_, found := library.members[memberName]
	if !found {
		library.members[memberName] = make(map[string]int)

		library.historyRecords = append(
			library.historyRecords,
			HistoryRecord{
				memberName:  memberName,
				bookName:    NA,
				description: "registered: a member",
				timeStamp:   time.Now()})
	}
}

func addBook(library *Library, bookName string) {
	_, found := library.books[bookName]

	if !found {
		library.books[bookName] = CopiesRecord{total: 1, lent: 0, remaining: 1}
	} else {
		library.books[bookName] = CopiesRecord{
			total:     library.books[bookName].total + 1,
			lent:      library.books[bookName].lent,
			remaining: library.books[bookName].remaining + 1}
	}

	library.historyRecords = append(
		library.historyRecords,
		HistoryRecord{
			memberName:  library.name,
			bookName:    bookName,
			description: "added: a copy",
			timeStamp:   time.Now()})
}

func validateBook(library *Library, bookName string) bool {
	_, found := library.books[bookName]

	if found {
		return true
	}

	fmt.Println(library.name, "validated that the book:", bookName, "does not exist")
	library.historyRecords = append(
		library.historyRecords,
		HistoryRecord{
			memberName:  NA,
			bookName:    bookName,
			description: "validated: book does not exist in the library",
			timeStamp:   time.Now()})
	return false
}

func validateMember(library *Library, memberName string) bool {
	_, found := library.members[memberName]

	if found {
		return true
	}

	fmt.Println(library.name, "validated that the member:", memberName, "does not exist")
	library.historyRecords = append(library.historyRecords,
		HistoryRecord{
			memberName:  memberName,
			bookName:    NA,
			description: "validated: not a member",
			timeStamp:   time.Now()})
	return false
}

func checkout(library *Library, bookName string, memberName string) {

	if !validateBook(library, bookName) || !validateMember(library, memberName) {
		return
	}

	isNotAvailable := library.books[bookName].remaining < 1

	if isNotAvailable {
		fmt.Println("Member:", memberName, "wanted to check-out the book:", bookName, "which has no remaining copy")
		library.historyRecords = append(library.historyRecords,
			HistoryRecord{
				memberName:  memberName,
				bookName:    bookName,
				description: "attempted: a check-out but failed as there was no remaining copy",
				timeStamp:   time.Now()})
		return
	}

	library.members[memberName][bookName]++
	library.books[bookName] = CopiesRecord{
		total:     library.books[bookName].total,
		lent:      library.books[bookName].lent + 1,
		remaining: library.books[bookName].remaining - 1}
	timeStamp := time.Now()

	fmt.Println("Book:", bookName, "is successfully checked out by member:", memberName, "at time:", timeStamp)
	library.historyRecords = append(library.historyRecords,
		HistoryRecord{
			memberName:  memberName,
			bookName:    bookName,
			description: "succeeded: check-out of a copy",
			timeStamp:   timeStamp})

}

func checkin(library *Library, bookName string, memberName string) {
	if !validateBook(library, bookName) {
		return
	}

	record := library.books[bookName]
	copiesCheckedoutByMember := library.members[memberName][bookName]

	if (record.lent < 1) || (copiesCheckedoutByMember < 1) {
		fmt.Println("Book:", bookName, "has no copy checked out to member", memberName)
		fmt.Println("Member:", memberName, "wanted to check-in Book:", bookName, "which has no copy checked out by Member:", memberName)
		library.historyRecords = append(library.historyRecords,
			HistoryRecord{
				memberName:  memberName,
				bookName:    bookName,
				description: "attempted: a check-in but failed as there is no copy checked out by",
				timeStamp:   time.Now()})
		return
	}

	library.members[memberName][bookName]--
	library.books[bookName] = CopiesRecord{
		total:     library.books[bookName].total,
		lent:      library.books[bookName].lent - 1,
		remaining: library.books[bookName].remaining + 1}
	timeStamp := time.Now()

	fmt.Println("Member:", memberName, "sucessfully checked-in the book:", bookName, "at time:", timeStamp)
	library.historyRecords = append(library.historyRecords, HistoryRecord{
		memberName:  memberName,
		bookName:    bookName,
		description: "succeeded: check-in of a copy",
		timeStamp:   time.Now()})
}

func main() {
	library := Library{
		name: "Westmount Public",
		books: map[string]CopiesRecord{
			"Baby Rudin":     {total: 1, remaining: 1},
			"Papa Rudin":     {total: 1, remaining: 1},
			"Dummit & Foote": {total: 1, remaining: 1}},
		members: map[string]map[string]int{
			"Aanika": make(map[string]int),
			"Dan":    make(map[string]int)},
		historyRecords: make([]HistoryRecord, 0)}

	printStatus(&library)
	addBook(&library, "Munkres")
	addBook(&library, "Dummit & Foote")
	addMember(&library, "Anne")
	printStatus(&library)
	checkout(&library, "Baby Rudin", "Aanika")
	checkout(&library, "Baby Rudin", "Dan")
	checkout(&library, "Dummit & Foote", "Dan")
	checkout(&library, "Dummit & Foote", "Dan")
	printStatus(&library)
	checkin(&library, "Munkres", "Anne")
	checkin(&library, "Baby Rudin", "Aanika")
	printStatus(&library)
}
