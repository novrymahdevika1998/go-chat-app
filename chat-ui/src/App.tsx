import './App.css'
import {BrowserRouter as Router, Route, Routes} from "react-router-dom";
import CreateUser from "./CreateUser.tsx";
import MainChat from "./MainChat.tsx";
import Login from "./Login.tsx";

function App() {
  return (
    <Router>
        <Routes>
            <Route path={"/create-user"} element={<CreateUser />} />
            <Route path="/chat" element={<MainChat />} />
            <Route path="/chat/:channelId" element={<MainChat />} />
            <Route path="/" element={<Login />} />
        </Routes>
    </Router>
  )
}

export default App
