import React from 'react';
import Footer from './Footer';
import Navbar from './Navbar';
import SideNav from './SideNav';
import Today from './Today';
import logo from '../images/j-logo.jpg'
import CalendarComp from './CalendarComp';

const AccountPage = () => {
  return (
    <>
        <div className='account-page-container main'>
            <div className='header-container'>
                <img className='logo-small' src={logo} alt='logo'/>
                <div className='navbar-container'>
                    <Navbar />
                </div>
            </div>
            
            <div className='main-content flex'>
                <div className='sidenav-container'>
                     <SideNav />
                </div>                
                <Today />
                
                <CalendarComp />
                
            </div>
        </div>
        
        <Footer />
    </>
    
    
  )
}

export default AccountPage