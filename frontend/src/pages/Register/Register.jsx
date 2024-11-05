import React, { useState } from 'react'
import { Link, useNavigate } from 'react-router-dom'
import axios from 'axios';
import "./Register.css"

const Register = () => {

    const [username, setUsername] = useState("")
    const [password, setPassword] = useState("")
    const [email, setEmail] = useState("")
    const [error, setError] = useState(null);

    const navigate = useNavigate()

    const handleRegister = async (e) => {
        e.preventDefault();
        try {
            await axios.post("http://localhost:8080/register", { username, password, email });
            navigate("/login");
        } catch (err) {
            setError('Registration failed');
        }

    };

    return (
        <div className='register-container'>
            <h2 className='register-title'>Register</h2>
            <form className='register-form' onSubmit={handleRegister}>
                <input
                    type='text'
                    value={username}
                    placeholder='Username'
                    onChange={(e) => setUsername(e.target.value)}
                    required
                    className='register-input'
                />
                <input
                    type='password'
                    placeholder='Password'
                    value={password}
                    onChange={(e) => setPassword(e.target.value)}
                    required
                    className='register-input'

                />
                <input
                    type='email'
                    placeholder='Email'
                    value={email}
                    onChange={(e) => setEmail(e.target.value)}
                    required
                    className='register-input'

                />
                <button className='register-button' type='submit'>Register</button>
                {error && <span className='error-message'>{error}</span>}
            </form>
            <div className='login-link'>
            <span>Already have an account!</span>
            <Link to="/login">
                <button className='login-link-button'>Login</button>
            </Link>
            </div>
        </div>
    )
}

export default Register
