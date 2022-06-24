import React,{useState} from 'react';
import Calendar from 'react-calendar';
import './Calendar.css'
// import style from '../calendar-component/style.css'
// import calStyle from '../';
// import calScript from '../calendar-component/calendar';

const CalendarComp = () => {
  const [value, setValue] = useState(new Date());

  function onChange(nextValue){
    setValue(nextValue)
  }
 
  
  return (
  <>
   <div id='calendar'>
      <Calendar onChange={onChange} value={value}  />
    </div>
    
  </>
  )
}

export default CalendarComp;