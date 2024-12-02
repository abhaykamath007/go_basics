import React from 'react'
import HeroBanner from '../../../components/HeroBanner/HeroBanner'
import Statistics from '../../../components/admin/Statistics/Statistics'
import Footer from '../../../components/Footer/Footer'

const AdminDashboard = () => {
  return (
    <div>
        <HeroBanner/>
        <Statistics/>
        <Footer/>
    </div>
  )
}

export default AdminDashboard
