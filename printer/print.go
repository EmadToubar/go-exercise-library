package printer

func Menu() {

	print(
		"1. Register a new borrower.\n",
		"2. Add a new book to the catalog.\n",
		"3. Delete a borrower by last name.\n",
		"4. Delete a book by ISBN.\n",
		"5. List all borrowers\n",
		"6. List all books.\n",
		"7. List all loans (returned and unreturned loans.)\n",
		"8. Search for a borrower.\n",
		"9. Search for a book.\n",
		"10. Loan a book.\n",
		"11. Return a book.\n",
		"12. Quit.\n",
		"Please enter an option: ")
}
