import React, { useState, useEffect } from "react";
import api from '../api/api'
import { useNavigate } from 'react-router-dom';

export default function Listmentor() {
    const [mentors, setMentor] = useState([])
    const nav = useNavigate();

    const getMentor = async () => {
        try {
        await api.get('/user/mentor/mentorlist')
         .then((res) => {
            setMentor(res.data.data)
            console.log(res.data.data)
        })
        } catch (error) {
          console.log(error);
        }
    };
    
    useEffect(() => {
        getMentor();
    }, []);

    const filter = (e) => {
      const keyword = e.target.value;

      if (keyword !== '') {
        const filteredMentor = mentors.filter(mentor => {
          return mentor.name.toLowerCase().includes(keyword.toLowerCase());
        });
        setMentor(filteredMentor);
      } else {
        setMentor(mentors);
      }
    }
    

    const bookMentor = async (id) => {
      try {
        await api.get(`/user/booking/mentor/${id}`)
         .then((res) => {
            nav("/profile")
        })
        } catch (error) {
          console.log(error);
        }
    }

      return (

        <>

          <form className='bg-gray-200 p-5'>   
            <label className="mb-2 text-sm font-medium text-gray-900 sr-only">Search</label>
            <div className="relative mx-10 lg:mx-32">
                <div className="flex absolute inset-y-0 left-0 items-center pl-3 pointer-events-none">
                    <svg className="w-5 h-5 text-gray-500" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"></path></svg>
                </div>
                <input type="search" 
                onChange={filter}
                className="block p-4 pl-10 w-full text-sm text-gray-900 bg-gray-50 rounded-lg border border-gray-300 focus:ring-yellow-500 focus:border-yellow-500" placeholder="Search Mentor..." required />
            </div>
          </form>


        <div className="bg-gray-200">
        <h1 className="flex justify-center text-4xl md:text-5xl lg:text-5xl font-bold">Daftar Mentor</h1>
        <div className=" bg-gray-200 flex justify-center flex flex-wrap items-center">
        {mentors && mentors.length > 0 ? (
          mentors.map((mentor, index) => {
            return (
              <div className="container max-w-lg bg-white rounded hadow-lg transform duration-200 easy-in-out m-20 hover:scale-110" key={index}>
              <div className="flex flex-wrap h-2/4 sm:h-64 overflow-hidden">
                  <img
                    className="w-full rounded-t"
                    src="https://ik.imagekit.io/t3defxmi6/hero_-aa6cXoCT.png?ik-sdk-version=javascript-1.4.3&updatedAt=1654854480679"
                    alt="bg"
                  />
              </div>
                <div className="flex justify-start px-5 -mt-12 mb-5">
                  <span clspanss="block relative h-32 w-32">
                    <img
                      src={mentor.image}
                      alt="avatar"
                      className="mx-auto object-cover rounded-full h-24 w-24 bg-white p-1"
                    />
                  </span>
                </div>
                <div className="">
                  <div className="px-8 mb-8">
                    <h2 className="mb-2 text-2xl font-bold tracking-tight text-gray-900">
                    {mentor.name}
                    </h2>
                    <p className="text-yellow-400 mt-2">
                    Mentor {mentor.skill}
                    </p>
                    <p className="mt-2 text-gray-500">
                      {mentor.bio}
                    </p>
                    <button
                    onClick={() => bookMentor(mentor.id)}
                    className="mt-5 inline-flex items-center py-2 px-3 text-sm font-medium text-center text-white bg-yellow-500 rounded-lg hover:bg-yellow-700 focus:ring-4 focus:outline-none focus:ring-yellow-200">
                      Buat Janji
                    </button>
                  </div>
                </div>
              </div> 
         
            )
          }

          )
        ) : ( <div className="text-center">
          <h1 className="text-3xl text-gray-400 font-bold">Tidak ada mentor</h1>
        </div> )}

        </div>
        </div>

        </>
      )
    }    

