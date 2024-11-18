import React, { useEffect, useState } from 'react'
import BookItem from '../../components/BookItem/BookItem';
import axiosInstance from '../../utils/axiosInstance';
import "./BooksPage.css"

const BooksPage = () => {
  const [books, setBooks] = useState([]);
  const genres = ["Fantasy", "Fiction", "Romance", "Adventure"]
  const [selectedGenre, setSelectedGenre] = useState('');
  const [selectedStatus, setSelectedStatus] = useState('');
  const [filteredBooks, setFilteredBooks] = useState([]);

  const [currentPage, setCurrentPage] = useState(1);
  const booksPerPage = 10;
  const indexOfLastBook = currentPage * booksPerPage;
  const indexOfFirstBook = indexOfLastBook - booksPerPage;
  const currentBooks = filteredBooks.slice(indexOfFirstBook, indexOfLastBook);
  const totalPages = Math.ceil(filteredBooks.length / booksPerPage);

  useEffect(() => {
    const fetchBooks = async () => {
      try {
        const response = await axiosInstance.get('/books');
        setBooks(response.data);
        setFilteredBooks(response.data);
      } catch (err) {
        console.error("Error fetching books:", err);
      }
    };
    fetchBooks();
  }, [])

  useEffect(() => {
    let updatedBooks = books;

    if (selectedGenre) {
      updatedBooks = updatedBooks.filter((book) => book.genre === selectedGenre);
    }

    if (selectedStatus) {
      updatedBooks = updatedBooks.filter((book) => book.availability_status === selectedStatus);
    }

    setFilteredBooks(updatedBooks);
    setCurrentPage(1);
  }, [selectedGenre, selectedStatus, books]);


  const handleGenreChange = (event) => {
    setSelectedGenre(event.target.value);
  }

  const handleStatusChange = (event) => {
    setSelectedStatus(event.target.value);
  }


  const goToNextPage = () => {
    if (currentPage < totalPages)
      setCurrentPage(prePage => prePage + 1)
  }

  const goToPreviousPage = () => {
    if (currentPage > 1)
      setCurrentPage(prePage => prePage - 1)
  }

  return (
    <div className='books-page'>
      <div className='books-page-heading'>
        <h1>Library Books</h1>
        <div className='filter-section'>
          <div className='genre-dropdown'>
          <label htmlFor='genre'>Filter by Genre:</label>
          <select id="genre" value={selectedGenre} onChange={handleGenreChange}>
            <option value="">All Genres</option>
            {
              genres.map((genre) => (
                <option key={genre} value={genre}>{genre}</option>
              ))
            }
          </select>
          </div>
          <div className='status-dropdown'>
          <label htmlFor='status'>Book status:</label>
          <select id='status' value={selectedStatus} onChange={handleStatusChange}>
            <option value="">All Books</option>
            <option value="available">Available</option>
            <option value="checked_out">Checked Out</option>
          </select>
          </div>
        </div>
      </div>
      <div className='books-grid'>
        {currentBooks.length > 0 ? (
          currentBooks.map((book) => (
            <BookItem key={book.id} book={book} />
          ))
        ) : (
          <p className='book-not-found'>No books found... </p>
        )}
      </div>
      {filteredBooks.length > 0 &&
        <div className='pagination'>
          <button onClick={goToPreviousPage} disabled={currentPage === 1}>
            Previous
          </button>
          <span>Page {currentPage} of {totalPages}</span>
          <button onClick={goToNextPage} disabled={currentPage === totalPages}>
            Next
          </button>
        </div>
      }
    </div>
  )
}

export default BooksPage
