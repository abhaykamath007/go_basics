import React from "react";
import Login from "./pages/Login/Login";
import { BrowserRouter as Router, Route, Routes} from 'react-router-dom';
import Register from "./pages/Register/Register";

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/login" element={<Login />} />
        <Route path="/register" element={<Register />} />
      </Routes>
    </Router>
  );
}

export default App;
