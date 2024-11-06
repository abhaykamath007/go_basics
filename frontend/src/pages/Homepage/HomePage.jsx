import React from 'react'
import Navbar from '../../components/Navbar/Navbar'
import HeroBanner from '../../components/HeroBanner/HeroBanner'
import FeaturedBooks from '../../components/FeaturedBooks/FeaturedBooks'

const HomePage = () => {
  return (
    <div>
      <Navbar/>
      <HeroBanner/>
      <FeaturedBooks/>
    </div>
  )
}

export default HomePage
