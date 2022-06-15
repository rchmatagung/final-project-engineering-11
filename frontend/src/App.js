import SignIn from './pages/SignIn';
import SignUp from "./pages/SignUp";
import Navbar from "./components/Navbar";
import Hero from './components/Hero';
import Articles from './components/Articles';
import AboutUsHero from "./components/AboutUsHero"
import OurTeam from "./components/OurTeam";
import Footer from "./components/Footer";

function App() {
  return (
    <div className="App">
      {/* <SignIn /> */}
      {/* <SignUp /> */}

      <Navbar />
      <Hero />
      <Articles />
      {/* <AboutUsHero />
      <OurTeam /> */}
      <Footer />
    </div>
  );
}

export default App;
