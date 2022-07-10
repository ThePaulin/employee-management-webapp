import React, { useState } from 'react';
import clsx from 'clsx'
import Footer from '../Footer';
import logo from '../../images/j-logo.jpg';
import { useLoginFormValidator } from './Hooks/useLoginValidator';

const LOGIN = {
    username: "username",
    password: "password",
    admin: "admin"
}

const Login = (props) => {
    const [form, setForm] = useState({
        [LOGIN.username]: "",
        [LOGIN.password]:"",
        [LOGIN.admin]: false
    });


//passing the form to validator to get error messages
    const {errors, validateForm, onBlurField} = useLoginFormValidator(form)

   




    

    const handleChange = e=>{
            const field = e.target.name;
            const nextFormState = {
                ...form, 
                [field]: e.target.value
            };
            setForm(nextFormState);
            if(errors[field].dirty){
                validateForm({
                    form: nextFormState,
                    errors,
                    field
                });
            };
            // console.log(form)
    
    };
    const handleCheckbox = e=>{
        setForm(prev=>{
            const field = e.target.name
            return {
                ...form,
                [field]: !form.admin
            }
        });
    };

    const handleSubmit = e=>{
        e.preventDefault();
        const {isValid} = validateForm({form, errors, forceTouchErrors: true});
        if(!isValid) return;
        alert(JSON.stringify(form,null,2));
    };
  return (
    <>
        <div className='login main'>
            <div className='login-container flex'>
                <img className='logo-login' src={logo} alt="logo" />
                <form className='login-form flex' onSubmit={handleSubmit}>
                    <div className='login-message'>Login To Access</div>
                    <input 
                        // className={clsx(
                        //     errors.email.dirty &&
                        //     errors.email.error &&
                        //     "formFieldError"
                        // )}
                        type='text' 
                        name= {LOGIN.username} 
                        onChange={handleChange} 
                        onBlur={onBlurField}
                        placeholder='Enter Username' 
                        required={true}
                    />
                    {errors.username.dirty && errors.username.error ? (
                        <p className="formFieldErrorMessage">{errors.username.message}</p>
                    ): null}
                    <input 
                        // className={clsx(
                        //     errors.password.dirty &&
                        //     errors.password.error &&
                        //     "formFieldError"
                        // )}
                        type='password' 
                        name={LOGIN.password} 
                        onChange={handleChange} 
                        onTouchEnd={handleChange} 
                        onBlur={onBlurField}
                        placeholder='Enter Password' 
                        required={true} 
                    />
                    {errors.password.dirty && errors.password.message ? (
                        <p className='formFieldErrorMessage'>{errors.password.message}</p>
                    ): null}
                    <div className='admin-div flex'>
                        <label id='admin-checkbox-label' for='admin-checkbox'> Admin Access</label>
                        <input 
                            id='admin-checkbox' 
                            name={LOGIN.admin} 
                            type='checkbox' 
                            onChange={handleCheckbox} 
                        />                        
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