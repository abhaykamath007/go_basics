import React from 'react'
import "../BorrowedBooks/BorrowedBooks.css"

const ReturnedBooks = ({books}) => {
  return (
    <div className="borrowed-books">
      <h2>Returned Books</h2>
      {books && books.length > 0 ? (
        <table className="borrowed-books-table">
          <thead>
            <tr>
              <th>Title</th>
              <th>Author</th>
              <th>Borrowed Date</th>
              <th>Due Date</th>
            </tr>
          </thead>
          <tbody>
            {books.map((book) => (
              <tr key={book.id}>
                <td>{book.title}</td>
                <td>{book.author}</td>
                <td>{new Date(book.borrowed_date).toLocaleDateString()}</td>
                <td>{new Date(book.due_date).toLocaleDateString()}</td>
              </tr>
            ))}
          </tbody>
        </table>
      ) : (
        <p>No Returned books found.</p>
      )}
    </div>
  );
}

export default ReturnedBooks
