import React, { useEffect, useState } from 'react'
import BookItem from '../../components/BookItem/BookItem';
import axiosInstance from '../../utils/axiosInstance';
import "./BooksPage.css"

const BooksPage = () => {
  const [books, setBooks] = useState([]);
  const [genres, setGenres] = useState([]);
  const [selectedGenre, setSelectedGenre] = useState("");
  const [selectedStatus, setSelectedStatus] = useState("");

  const [currentPage, setCurrentPage] = useState(1);
  const [totalPages, setTotalPages] = useState(1);

  useEffect(()=> {
    const fetchGenres = async () => {
      try {
        const response = await axiosInstance.get("/genres");
        setGenres(response.data.genres);
      } catch(err) {
        console.error("Error fetching genres : ",err);
      }
    };
    fetchGenres();
  },[]);


  useEffect(() => {
    const fetchBooks = async () => {
      try {
        const response = await axiosInstance.get(`/books`,{
          params: {
            page: currentPage,
            genre: selectedGenre || undefined,
            status: selectedStatus || undefined,
          },
        });
        const booksData = response.data.books || []; 
        setBooks(booksData);
        setTotalPages(response.data.totalPages || 1);
      } catch (err) {
        console.error("Error fetching books:", err);
        setBooks([]);
        setTotalPages(1);
      }
    };
    fetchBooks();
  }, [currentPage,selectedGenre,selectedStatus]);

  const handleGenreChange = (event) => {
    setSelectedGenre(event.target.value);
    setCurrentPage(1);
  }

  const handleStatusChange = (event) => {
    setSelectedStatus(event.target.value);
    setCurrentPage(1);
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
        {books.length > 0 ? (
          books.map((book) => (
            <BookItem key={book.id} book={book} />
          ))
        ) : (
          <p className='book-not-found'>No books found... </p>
        )}
      </div>
      {books.length > 0 &&
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
