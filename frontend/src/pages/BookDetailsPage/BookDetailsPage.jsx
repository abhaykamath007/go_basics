import React, { useEffect, useState } from 'react'
import { useNavigate, useParams } from 'react-router-dom'
import axiosInstance from '../../utils/axiosInstance';
import NotFoundPage from '../NotFoundPage/NotFoundPage';
import PosterFallbacks from '../../assets/no-poster.png'
import { toast } from "react-toastify";
import "./BookDetailsPage.css"
import { useAuth } from '../../context/AuthContext';

const BookDetailsPage = () => {
    const { id } = useParams();
    const [book, setBook] = useState(null)
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);
    const [isBorrowed, setIsBorrowed] = useState(false);
    const navigate = useNavigate();
    const { userID } = useAuth();
    useEffect(() => {
        const fetchBookDetails = async () => {
            try {
                const response = await axiosInstance.get(`/books/${id}`);
                setBook(response.data);
                setIsBorrowed(response.data.availability_status !== 'available');
                setLoading(false);
            } catch (err) {
                setError("Failed to fetch book details");
                setLoading(false);
            }
        };

        fetchBookDetails();
    }, [id])

    const handleBorrow = async()=>{
        try {
            const response = await axiosInstance.post(`/books/${id}/borrow`,{
                user_id:userID,
            });
      
            if (response.status === 200) {
              toast.success("You have borrowed the book successfully!");
              setIsBorrowed(true);  
            }
          } catch (error) {
            toast.error("Failed to borrow the book. Please try again.");
          }
    }

    if (loading) {
        return <div>Loading...</div>
    }

    if (error) {
        return <NotFoundPage />
    }

    if (!book) {
        return <div>No book found</div>
    }
    return (
        <div className="book-details-container">
            <div className="book-image">
                <img src={PosterFallbacks} alt={book.title} />
            </div>
            <div className="book-details">
                <h1>{book.title}</h1>
                <p><strong>Author : </strong> {book.author}</p>
                <p><strong>Genre :</strong> {book.genre}</p>
                <p><strong>Year :</strong> {book.publication_year}</p>
                <div className='description'> 
                <strong>Description:</strong> Lorem ipsum dolor sit amet consectetur adipisicing elit. Facilis at sint in ad voluptate impedit consectetur enim obcaecati minima qui? Fugit facilis similique blanditiis provident accusamus perferendis quibusdam, temporibus qui.
                </div>
                <p><strong>Status: </strong>{book.availability_status}</p>
                <button onClick={handleBorrow} disabled={isBorrowed} className={`borrow-btn ${isBorrowed} ?"disabled" : ""`}> {isBorrowed ? 'Not Available' : 'Borrow'}</button>
            </div>
        </div>
    )
}

export default BookDetailsPage
