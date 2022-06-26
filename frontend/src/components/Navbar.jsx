import React, { useState } from 'react'
import logoImg from '../assets/logohi.png'
import {MenuIcon, XIcon} from '@heroicons/react/outline'
import { Link, useNavigate } from 'react-router-dom';
import api from '../api/api'

export default function Navbar() {
  const [navbarOpen, setNavbarOpen] = useState(false);
  const token = localStorage.getItem('token');
  const nav = useNavigate();

  const logout = async () => {
      try {
        await api.post('/auth/logout')
        .then((res) => {
          console.log(res)
          localStorage.removeItem('token');
          localStorage.removeItem('RLPP');
          localStorage.removeItem('id');
          nav("/")
      })
      } catch (error) {
        console.log(error);
      }
  };

  return (
    <>
      <nav className="relative
          w-full
          flex flex-wrap
          items-center
          justify-between
          py-4
          bg-gray-100
          text-gray-500
          hover:text-gray-700
          focus:text-gray-700
          shadow-lg
          navbar navbar-expand-lg navbar-light
          ">
        <div className="container-fluid w-full flex flex-wrap items-center justify-between px-6">
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
                <Link to={`/`} > Home </Link>
              </li>
              <li className="px-5 py-2 flex items-center text-lg hover:font-bold">
                <Link to={`/about`} > About </Link>
              </li>
              <li className="px-5 py-2 flex items-center text-lg hover:font-bold">
                <Link to={`/findmentor`} > Find Mentor </Link>
              </li>
              <br/>
            </ul>

            {token !== null ? (
              <div className="flex flex-col md:flex-row md:justify-end">
                <Link to={`/profile`} className="px-5 py-3 hover:font-bold">Profile</Link>
                <button onClick={logout} className="rounded-lg bg-yellow-500 px-5 py-3 hover:font-bold">Sign Out</button>
              </div>
            ) : (
              <div className="flex flex-col md:flex-row md:justify-end">
                <Link to={`/signup`} className="rounded-lg bg-yellow-500 px-5 py-3 hover:font-bold">Sign Up</Link>
                <Link to={`/signin`} className="px-5 py-3 hover:font-bold">Sign In</Link>
              </div>
            )}
          </div>
        </div>
      </nav>
    </>
  );
}