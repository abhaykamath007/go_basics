import axios from "axios";
import { createContext, useContext, useState } from "react"
import { useNavigate } from "react-router-dom";

const AuthContext = createContext();

export const AuthProvider = ({children}) => {
    const [isAuthenticated,setIsAuthenticated] = useState(false);
    const [loading,setLoading] = useState(false);
    const [username,setUsername] = useState(null);
    const navigate = useNavigate();

    const login = async (username,password) => {
        setLoading(true);
        try {
            const response = await axios.post("http://localhost:8080/login",{username,password})

            const {token,expirationTime} = response.data;
            localStorage.setItem("token",token);
            localStorage.setItem("expirationTime",expirationTime);

            setUsername(username);
            setIsAuthenticated(true);
            
        } catch(err) {
            throw new Error("Login failed. Please try again");
        } finally {
            setLoading(false);
        }
    }

    const logout = () => {
        localStorage.removeItem("token");
        localStorage.removeItem("expirationTime");
        setIsAuthenticated(false);
        navigate("/login");
    }

    return (
        <AuthContext.Provider value={{isAuthenticated,login,logout,loading,username}}>
        {children}
        </AuthContext.Provider>
        
    )
}

export const useAuth = () => useContext(AuthContext)