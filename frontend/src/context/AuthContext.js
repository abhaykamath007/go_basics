import axios from "axios";
import { createContext, useContext, useEffect, useState } from "react"
import { useNavigate } from "react-router-dom";

const AuthContext = createContext();

export const AuthProvider = ({children}) => {
    const [isAuthenticated,setIsAuthenticated] = useState(false);
    const [loading,setLoading] = useState(false);
    const [username,setUsername] = useState(null);
    const [userID,setUserID] = useState(null);
    const navigate = useNavigate();


    useEffect(()=>{
        const token = localStorage.getItem('token');
        const user = localStorage.getItem('username');
        const authStatus = localStorage.getItem("isAuthenticated") === "true"; 
        setUsername(user);

        if(token){
            setIsAuthenticated(true);
        }
    },[])

    const login = async (username,password) => {
        setLoading(true);
        try {
            const response = await axios.post("http://localhost:8080/login",{username,password})

            const {token,expirationTime,UserID} = response.data;
            console.log(UserID);
            localStorage.setItem("token",token);
            localStorage.setItem("expirationTime",expirationTime);
            localStorage.setItem("username",username);
            localStorage.setItem("isAuthenticated", "true");
            localStorage.setItem("userID",UserID);
            setUsername(username);
            setUserID(UserID);
            setIsAuthenticated(true);
            navigate("/");

        } catch(err) {
            throw new Error("Login failed. Please try again");
        } finally {
            setLoading(false);
        }
    }

    const logout = () => {
        localStorage.removeItem("token");
        localStorage.removeItem("expirationTime");
        localStorage.removeItem("isAuthenticated");
        localStorage.removeItem("username");
        localStorage.removeItem("userID");
        setUsername(null);
        setIsAuthenticated(false);
        navigate("/login");
    }

    return (
        <AuthContext.Provider value={{isAuthenticated,login,logout,loading,username,userID}}>
        {children}
        </AuthContext.Provider>
        
    )
}

export const useAuth = () => useContext(AuthContext)


