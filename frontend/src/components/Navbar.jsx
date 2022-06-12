import React, {useState} from 'react'
import logoImg from '../assets/logohi.png'
import {MenuIcon, XIcon} from '@heroicons/react/outline'

export default function Navbar() {
  const [navbarOpen, setNavbarOpen] = useState(false);

  return (
    <>
      <nav className="relative flex flex-wrap items-center justify-between bg-transparent px-2 py-3 mb-3">
        <div className="container px-4 mx-auto flex flex-wrap items-center justify-between">
          <div className="w-full relative flex justify-between md:w-auto md:static md:block md:justify-start">
            <div> 
              <img className='h-8 w-15' src={logoImg} alt="Hicoder" />
            </div>
            <button
              className="md:hidden"
              type="button"
              onClick={() => setNavbarOpen(!navbarOpen)}>
              {(!navbarOpen) ? <MenuIcon className="w-8" /> : <XIcon className="w-8"/>}
            </button>
          </div>
          <div
            className={"md:flex flex-grow items-center" +
              (navbarOpen ? " flex" : " hidden")
            }
          >
            <ul className="flex flex-col md:flex-row list-none mr-auto">
              <li className="px-5 py-2 flex items-center text-lg hover:font-bold">
                <a href="/" > Home </a>
              </li>
              <li className="px-5 py-2 flex items-center text-lg hover:font-bold">
                <a href="/"> Article </a>
              </li>
              <li className="px-5 py-2 flex items-center text-lg hover:font-bold">
                <a href="/"> Find Mentor </a>
              </li>
              <li className="px-5 py-2 flex items-center text-lg hover:font-bold">
                <a href="/"> Course </a>
              </li>
              <br/>
            </ul>
            <div className="flex flex-col md:flex-row md:justify-end">
                 <button className="rounded-lg bg-yellow-500 px-5 py-3"> Sign In </button>
                 <button className="px-5 py-3"> Sign Up </button>
             </div>
          </div>
        </div>
      </nav>
    </>
  );
}