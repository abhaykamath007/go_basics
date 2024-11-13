
import React from "react";
import "./BorrowedBooks.css";
import axiosInstance from "../../utils/axiosInstance";
import { toast } from "react-toastify";
import { useAuth } from "../../context/AuthContext";

const BorrowedBooks = ({ books }) => {

  const { userID } = useAuth();
  const handleReturnBook = async (bookId) => {
    try {
      const response = await axiosInstance.post(`/books/${bookId}/return`,{
        user_id:userID,
    });
      if (response.status == 200) {
        toast.success("Book returned successfully!");
      } 
    } catch (error) {
      toast.error("Failed to return book. Please try again");
    }
  };

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
            {books.map((book) => (
              <tr key={book.id}>
                <td>{book.title}</td>
                <td>{book.author}</td>
                <td>{new Date(book.borrowed_date).toLocaleDateString()}</td>
                <td>{new Date(book.due_date).toLocaleDateString()}</td>
                <td><button onClick={()=> handleReturnBook(book.id)}>Return</button></td>
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
