import React from 'react';
import Footer from './Footer';
import logo from '../images/j-logo.jpg';

const Login = () => {
  return (
    <>
        <div className='login main'>
            <div className='login-container flex'>
                <img className='logo-login' src={logo} alt="logo" />
                <form className='login-form flex'>
                    <div className='login-message'>Login To Access</div>
                    <input type='text' placeholder='Enter Username'/>
                    <input type='text' placeholder='Enter Password' />
                    <div className='admin-div flex'>
                        <label id='admin-checkbox-label' for='admin-checkbox'> Admin Access</label>
                        <input id='admin-checkbox' type='checkbox' />
                        
                    </div> 
                    <button type='submit'>Login</button>
                </form>
            </div>
        </div>
        <Footer />
        
    </>
    
    
  )
}

export default Login