import "./FeaturedBooks.css"
import React, { useEffect, useState } from 'react'
import axiosInstance from '../../utils/axiosInstance';

const FeaturedBooks = () => {

    const [books,setBooks] = useState([]);

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

  return (
    <div className='book-list'>
      <h2>Featured Books</h2>
      <div className='book-grid'>
        {books.length > 0 ? (
            books.map(book => (
                <div className='book-item' key={book.id}>
                    <img src={'http://i2.wp.com/geekdad.com/wp-content/uploads/2013/02/HP1-Kibuishi.jpg'} alt={book?.title} className='book-image'/>
                    <h3 className='book-title'>{book.title}</h3>
                    <p className='book-author'>{book.author}</p>
                    <button className='book-button'>Borrow</button>
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
