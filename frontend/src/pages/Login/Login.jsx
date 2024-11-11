import React, { useState } from 'react'
import { Link, useNavigate } from 'react-router-dom';
import "./Login.css"
import  { useAuth } from '../../context/AuthContext';

const Login = () => {

    const {login,loading} = useAuth();
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    const [error, setError] = useState(null);
    const navigate = useNavigate();

    const handleLogin = async (e) => {
        e.preventDefault();
        try {
            await login(username,password);
        } catch (err) {
            setError("Login failed. Please try again");
        }
    }

    return (
        <div className="login-container">
            <h2 className="login-title">Login</h2>
            <form onSubmit={handleLogin} className="login-form">
                <input
                    type="text"
                    value={username}
                    onChange={(e) => setUsername(e.target.value)}
                    placeholder="Username"
                    className="login-input"
                    required
                />
                <input
                    type="password"
                    value={password}
                    onChange={(e) => setPassword(e.target.value)}
                    placeholder="Password"
                    className="login-input"
                    required
                />
                <button type="submit" className="login-button">Login</button>
                {error && <span className="error-message">{error}</span>}
            </form>
            <div className="register-link">
                <span>Don't have an account?</span>
                <Link to="/register">
                    <button className="register-link-button">Register</button>
                </Link>
            </div>
        </div>
    )
}

export default Login
