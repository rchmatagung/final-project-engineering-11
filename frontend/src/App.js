import SignIn from './pages/SignIn';
import SignUp from "./pages/SignUp";
import Navbar from "./components/Navbar";
import Hero from './components/Hero';
import Articles from './components/Articles';
import AboutUsHero from "./components/AboutUsHero"
import OurTeam from "./components/OurTeam";
import Footer from "./components/Footer";
import {Routes, Route} from 'react-router-dom';
import ProfileCard from './components/ProfileCard';
import BookingStatus from './components/BookingStatus'
import Listmentor from './components/Listmentor';
import Search from './components/Search';

function App() {
  return (
    <div className="App">
      <Routes>
        <Route path="/" element={
          <div>
            <Navbar />
            <Hero />
            <Articles />
            <Footer />
          </div>
        } />
        <Route path="/about" element={
          <div>
            <Navbar />
            <AboutUsHero />
            <OurTeam />
            <Footer />
          </div>
        }/>
        <Route path="/findmentor" element={
          <div>
            <Navbar />
            <Search/>
            <Listmentor/>
            <Footer/>
          </div>
        }/>
        <Route path="/signin" element={<SignIn />} />
        <Route path="/signup" element={<SignUp />} />
        <Route path="/profile" element={
          <div>
            <Navbar />
            <ProfileCard />
            <BookingStatus />
            <Footer />
          </div>
        }/>
      </Routes>
    </div>
  );
}

export default App;
