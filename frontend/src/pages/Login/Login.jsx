import React, { useState } from 'react'
import { Link } from 'react-router-dom';

const Login = () => {

    const [username,setUsername] = useState("");
    const [password,setPassword] = useState("");

    const handleLogin = (e) => {
        e.preventDefault();

    }

    return (
        <div>
            <h2>Login</h2>
            <form onSubmit={handleLogin}>
                <input
                    type='text'
                    value={username}
                    onChange={(e) => setUsername(e.target.value)}
                    placeholder="Username"
                />
                <input
                    type="password"
                    value={password}
                    onChange={(e)=> setPassword(e.target.value)}
                    placeholder="Password"
                />
                <button type="submit">Login</button>
                <span>Don't you have an account?</span>
                <Link to="/register">
                    <button>Register</button>
                </Link>
            </form>
        </div>
    )
}

export default Login
