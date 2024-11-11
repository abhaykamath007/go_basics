import React, { useState } from 'react'
import BorrowedBooks from '../../components/BorrowedBooks/BorrowedBooks'
import ReturnedBooks from '../../components/ReturnedBooks/ReturnedBooks'
import "./AccountPage.css"

const AccountPage = () => {
  const [activeSection,setActiveSection] = useState("borrowed");
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
        {activeSection === "borrowed" ? <BorrowedBooks/> : <ReturnedBooks/>}
      </main>
    </div>

  )
}

export default AccountPage
