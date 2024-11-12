
import React from "react";
import PropTypes from "prop-types";
import "./BorrowedBooks.css";

const BorrowedBooks = ({ books }) => {
  return (
    <div className="borrowed-books">
      <h2>Borrowed Books</h2>
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
            {books.map((book, index) => (
              <tr key={book.id || index}>
                <td>{book.title}</td>
                <td>{book.author}</td>
                <td>{new Date(book.borrowed_date).toLocaleDateString()}</td>
                <td>{new Date(book.due_date).toLocaleDateString()}</td>
              </tr>
            ))}
          </tbody>
        </table>
      ) : (
        <p>No borrowed books found.</p>
      )}
    </div>
  );
};


export default BorrowedBooks;
