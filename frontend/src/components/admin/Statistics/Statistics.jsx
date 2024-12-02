import React, { useEffect, useState } from 'react'
import "./Statistics.css"
import axiosInstance from '../../../utils/axiosInstance';

const Statistics = () => {
    const [stats, setStats] = useState({
        totalUsers: 0,
        totalBooks: 0,
        borrowedBooks: 0,
        availableBooks: 0,
        booksAddedThisMonth: 0,
        booksBorrowedThisMonth: 0,
    });

    const fetchStats = async () => {
        try {
            const response = await axiosInstance.get(`/admin/stats`);
            console.log(response.data);
            setStats({
                totalUsers: response.data.total_users,
                totalBooks: response.data.borrowed_books,
                availableBooks: response.data.available_books,
                borrowedBooks: response.data.borrowed_books,
                booksAddedThisMonth: response.data.books_added_this_month,
                booksBorrowedThisMonth: response.data.books_borrowed_this_month,
            });
        } catch (error) {
            console.error("Error fetching stats:", error);
        }
    }
    useEffect(() => {
        fetchStats();
    }, [])

    return (
            <div className='statistics-container'>
                <div className='stats-card'>
                    <h3>Total Users : </h3>
                    <p>{stats.totalUsers}</p>
                </div>
                <div className='stats-card'>
                    <h3>Total Books : </h3>
                    <p>{stats.totalBooks}</p>
                </div>
                <div className='stats-card'>
                    <h3>Borrowed Books : </h3>
                    <p>{stats.borrowedBooks}</p>
                </div>
                <div className='stats-card'>
                    <h3>Available Books : </h3>
                    <p>{stats.availableBooks}</p>
                </div>
                <div className='stats-card'>
                    <h3>Books Added this month : </h3>
                    <p>{stats.booksAddedThisMonth}</p>
                </div>
                <div className='stats-card'>
                    <h3>Books Borrowed this month : </h3>
                    <p>{stats.booksBorrowedThisMonth}</p>
                </div>
            </div>
    )
}

export default Statistics
