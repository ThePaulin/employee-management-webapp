import React from 'react'

const Footer = () => {
    const currentYear = new Date().getFullYear();
   
  return (
    <div className='footer'>
        <footer>
            <p>Copyright AGM Inc. {currentYear}</p>
        </footer>
    </div>
  )
}

export default Footer