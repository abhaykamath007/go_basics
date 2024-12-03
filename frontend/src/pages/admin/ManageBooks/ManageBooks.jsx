import React, { useEffect, useState } from 'react'
import { MdModeEdit } from "react-icons/md";
import { MdDelete } from "react-icons/md";
import EditBookModal from '../../../components/admin/EditBookModal/EditBookModal';
import DeleteBookModal from '../../../components/admin/DeleteBookModal/DeleteBookModal';
import axiosInstance from '../../../utils/axiosInstance';

const ManageBooks = () => {

    const [books, setBooks] = useState([]);
    const [editBook, setEditBook] = useState(null);
    const [deleteBook, setDeleteBook] = useState(null);

    const fetchBooks = async () => {
        try {
            const response = await axiosInstance.get("/books");
            setBooks(response.data.books);
        } catch (err) {
            console.error("Error fetching books:", err);
        }
    }

    const handleDeleteBook = async (bookId) => {
        try{
            await axiosInstance.delete(`books/delete/${bookId}`);
            setDeleteBook(null);
        } catch (error) {
            console.error("Error deleting book:",error);
        }
    }

    const handleEditAndCreateBook = async (updatedBook) => {
        try{
            if(updatedBook.id){
                await axiosInstance.put(`admin/books/update/${updatedBook.id}`,updatedBook);
            } else {
                console.log(updatedBook);
                await axiosInstance.post(`admin/books/create`,updatedBook);
            }
            setEditBook(null);
        } catch (error) {
            console.error("Error updating book:",error);
        }
    }


    useEffect(() => {
        fetchBooks();
    }, []);

    return (
        <div className='manage-books-page'>
            <div className='manage-books-header'>
                <h2>Manage Books</h2>
                <button onClick={() => setEditBook({})} className='create-book-btn'>
                    Create Book
                </button>
            </div>
            <table className='books-table'>
                <thead>
                    <tr>
                        <th>Title</th>
                        <th>Author</th>
                        <th>Genre</th>
                        <th>Year</th>
                    </tr>
                </thead>
                <tbody>
                    {books.length > 0 && books?.map((book) => (
                        <tr key={book.id}>
                            <td>{book.title}</td>
                            <td>{book.author}</td>
                            <td>{book.genre}</td>
                            <td>{book.publication_year}</td>
                            <td>
                                <MdModeEdit onClick={() => setEditBook(book)} />
                                <MdDelete onClick={() => setDeleteBook(book)} />
                            </td>
                        </tr>
                    ))}
                </tbody>
            </table>

            {editBook !== null && (
                <EditBookModal
                    book={editBook}
                    onSave={handleEditAndCreateBook}
                    onCancel={()=>setEditBook(null)}
                />
            )}

            {deleteBook !== null && (
                <DeleteBookModal
                    book={deleteBook}
                    onConfirm={() => handleDeleteBook(deleteBook.id)}
                    onCancel={() => setDeleteBook(null)}
                />
            )}
        </div>
    )
}

export default ManageBooks
