import React from 'react'

const Footer = () => {
    const currentYear = new Date().getFullYear();

    function whileInView(){
      var elements;
      var windowHeight;

      function init(){
        elements = document.querySelectorAll('.footerText');
        console.log(elements);
        windowHeight = window.innerHeight;
      }
      function checkPosition(){
        for(var i = 0; i < elements.length; i++){
          var element = elements[i];
          var positionFromTop = elements[i].getBoundingClientRect().top;

          if (positionFromTop - windowHeight <= 0){
            element.classList.add('whileInView');
            
          }
        }
      }
      window.addEventListener('scroll', checkPosition);
      window.addEventListener('resize', init)
    }
    whileInView();
   
  return (
    <div className='footer'>
       
      <p className='footerText'>Copyright Jumpstart Inc. {currentYear}</p>
       
    </div>
  )
}

export default Footer