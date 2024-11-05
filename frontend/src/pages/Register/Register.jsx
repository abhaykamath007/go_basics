import React, { useState } from 'react'
import { Link } from 'react-router-dom'

const Register = () => {

    const [username, setUsername] = useState("")
    const [password, setPassword] = useState("")
    const [email, setEmail] = useState("")


    const handleRegister = (e) => {
        e.preventDefault();
    }

    return (
        <div>
            <h2>Register</h2>
            <form onSubmit={handleRegister}>
                <input
                    type='text'
                    value={username}
                    placeholder='Username'
                    onChange={(e) => setUsername(e.target.value)}
                />
                <input
                    type='password'
                    placeholder='Password'
                    value={password}
                    onChange={(e) => setPassword(e.target.value)}
                />
                <input
                    type='email'
                    placeholder='Email'
                    value={email}
                    onChange={(e) => setEmail(e.target.value)}
                />
            </form>
            <span>Already have an account!</span>
            <Link to="/login">
                <button>Login</button>
            </Link>
        </div>
    )
}

export default Register
