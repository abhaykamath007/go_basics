import "./FeaturedBooks.css"
import React, { useEffect, useState } from 'react'
import axiosInstance from '../../utils/axiosInstance';
import PosterFallbacks from '../../assets/no-poster.png'
import { useNavigate } from "react-router-dom";

const FeaturedBooks = () => {

    const [books,setBooks] = useState([]);
    const navigate = useNavigate();


    useEffect(()=> {

        const fetchFeaturedBooks = async () => {
            try {
                
                const response = await axiosInstance.get('/books');
                const limitedBooks = response.data.slice(0, 5);
                setBooks(limitedBooks);
            } catch (err) {
                console.error("Error fetching books:", err);
            }
        };
        fetchFeaturedBooks();

    },[])

    const handleBookClick = (id) => {
        navigate(`/books/${id}`);
    }

  return (
    <div className='book-list'>
      <h2>Featured Books</h2>
      <div className='book-grid'>
        {books.length > 0 ? (
            books.map(book => (
                <div className='book-item' key={book?.id} onClick={()=>handleBookClick(book.id)}>
                    <img src={PosterFallbacks} alt={book?.title} className='book-image'/>
                    <h3 className='book-title'>{book?.title.length > 27 ? `${book.title.substring(0,27)}...` :  book?.title}
                    </h3>
                    <p className='book-author'>{book?.author.length > 27 ? `${book.author.substring(0,27)}...` : book?.author}</p>
                    {/* <button className='book-button'>Borrow</button> */}
                </div>
            ))
        ) : (
            <p>Loading featued books...</p>
        )}
      </div>
    </div>
  )
}

export default FeaturedBooks
