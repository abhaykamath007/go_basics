// BookItem.jsx
import React from 'react';
import { useNavigate } from 'react-router-dom';
import PosterFallbacks from '../../assets/no-poster.png';
import './BookItem.css';

const BookItem = ({ book }) => {
    const navigate = useNavigate();

    const handleBookClick = () => {
        navigate(`/books/${book.id}`);
    };

    return (
        <div className='book-item' onClick={handleBookClick}>
            <img src={PosterFallbacks} alt={book?.title} className='book-image'/>
            <h3 className='book-title'>
                {book?.title.length > 27 ? `${book.title.substring(0, 27)}...` : book?.title}
            </h3>
            <p className='book-author'>
                {book?.author.length > 27 ? `${book.author.substring(0, 27)}...` : book?.author}
            </p>
        </div>
    );
};

export default BookItem;
