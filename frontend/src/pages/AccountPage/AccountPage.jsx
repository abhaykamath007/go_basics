import React, { useEffect, useState } from 'react'
import BorrowedBooks from '../../components/BorrowedBooks/BorrowedBooks'
import ReturnedBooks from '../../components/ReturnedBooks/ReturnedBooks'
import "./AccountPage.css"
import axiosInstance from '../../utils/axiosInstance'

const AccountPage = () => {
  const [activeSection, setActiveSection] = useState("borrowed");
  const [borrowedBooks, setBorrowedBooks] = useState([]);
  const [returnedBooks, setReturnedBooks] = useState([]);
  const [loading, setLoading] = useState(false);

  const userID = localStorage.getItem("userID")
  console.log(userID);

  const fetchBooks = async (section) => {
    setLoading(true);
    try {
      const endpoint = section === "borrowed" ? `users/books/${userID}/borrowed` : `users/books/${userID}/returned`;
      const response = await axiosInstance.get(endpoint);

      if (section === "borrowed") {
        setBorrowedBooks(response.data);
        console.log(borrowedBooks);
      } else {
        setReturnedBooks(response.data);
      }
    } catch (error) {
      console.error("Error fetching books:", error);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    
    fetchBooks(activeSection);
  }, [activeSection])
  
  const refreshBooks = () => {
    fetchBooks("borrowed");
  };


  return (
    <div className='account-page'>
      <aside className='sidebar'>
        <button
          className={`sidebar-option ${activeSection === "borrowed" ? "active" : ""}`}
          onClick={() => setActiveSection("borrowed")}
        >
          Borrowed Books
        </button>
        <button
          className={`sidebar-option ${activeSection === "returned" ? "active" : ""}`}
          onClick={() => setActiveSection("returned")}
        >
          Returned Books
        </button>
      </aside>
      <main className='account-content'>
        {loading ? (
          <p className='loading'>Loading....</p>
        ) : activeSection === "borrowed" ? (
          <BorrowedBooks books={borrowedBooks} onReturn={refreshBooks} />
        ) : (
          <ReturnedBooks books={returnedBooks} />
        )}
      </main>
    </div>

  )
}

export default AccountPage
