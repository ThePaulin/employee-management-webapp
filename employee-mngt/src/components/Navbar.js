import React from 'react';

const Navbar = () => {
    const name = "Paulin";
  return (
    <>
        <header className='flex'>
            <p id='nav-greeting'> Welcome back {name}!</p>
            <nav>
                <ul>
                    <li><a href='#'>ABOUT</a></li>
                    <li><a href='#'>HELP</a></li>
                </ul>
            </nav>
        </header>
    </>
  )
}

export default Navbar;