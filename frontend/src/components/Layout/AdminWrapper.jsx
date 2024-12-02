import React from 'react'
import { Outlet } from 'react-router-dom'
import Navbar from '../admin/Navbar/Navbar'

const AdminWrapper = () => {
  return (
    <>
      <Navbar/>
      <Outlet/>
    </>
  )
}

export default AdminWrapper
