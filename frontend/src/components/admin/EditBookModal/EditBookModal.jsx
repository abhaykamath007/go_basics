import React, { useState } from 'react'

const EditBookModal = ({book,onSave,onCancel}) => {
    const [formData,setFormData] = useState({
        title : book.title || '',
        author : book.author || '',
        genre : book.genre || '',
        publication_year : book.publication_year || '',
    })

    const handleSave = (e) => {
        e.preventDefault();
        onSave({...book,...formData})
    }

  return (
    <div className='modal'>
      <h3>{formData.id ? "Edit Book" : "Create Book"}</h3>
      <form onSubmit={handleSave}>
        <input
            type='text'
            value={formData.title}
            onChange={(e) => setFormData({ ...formData, title: e.target.value })}
            placeholder='Title'
            required
        />
        <input
            type='text'
            value={formData.author}
            onChange={(e) => setFormData({ ...formData, author: e.target.value })}
            placeholder='Author'
            required
        />
        <input
            type='text'
            value={formData.genre}
            onChange={(e) => setFormData({ ...formData, genre: e.target.value })}
            placeholder='Genre'
            required
        />
        <input
            type='text'
            value={formData.publication_year}
            onChange={(e) => setFormData({ ...formData, publication_year: parseInt(e.target.value) || "" })}
            placeholder='publication year'
            required
        />
        <div className='modal-actions'>
            <button type="submit">
                {book.id ? "Save" : "Create Book"}
            </button>
            <button type="button" onClick={onCancel}>
                Cancel
            </button>
        </div>
      </form>
    </div>
  )
}

export default EditBookModal
