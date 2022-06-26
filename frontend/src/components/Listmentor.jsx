import React, { useState, useEffect } from "react";
import api from '../api/api'

export default function Listmentor() {
    const [mentors, setMentor] = useState([])

    const getUser = async () => {
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
        getUser();
    }, []);

      return (
        <div className="bg-gray-200 dark:bg-white-500">
        <h1 className="flex justify-center  text-4xl md:text-5xl lg:text-5xl font-bold">Daftar Mentor</h1>
        <div className=" bg-gray-200 flex justify-center dark:bg-white-500 flex flex-wrap items-center">
          {mentors.map((mentor, index) => {
          return (
          <div className="container max-w-lg bg-white rounded dark:bg-gray-800 shadow-lg transform duration-200 easy-in-out m-20 hover:scale-110"  key={index}>
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
                        <h2 className="mb-2 text-2xl font-bold tracking-tight text-gray-900 dark:text-white">
                        {mentor.name}
                        </h2>
                        <p className="text-yellow-400 mt-2 dark:text-yellow-400">
                        Mentor {mentor.skill}
                        </p>
                        <p className="mt-2 text-gray-500 dark:text-gray-300">
                          {mentor.bio}
                        </p>
                        <button onClick={() => window.location = ''} className="mt-5 inline-flex items-center py-2 px-3 text-sm font-medium text-center text-white bg-yellow-500 rounded-lg hover:bg-yellow-700 focus:ring-4 focus:outline-none focus:ring-yellow-200 dark:bg-yellow-400 dark:hover:bg-yellow-600 dark:focus:ring-yellow-700">
                          Buat Janji
                        </button>
                      </div>
                    </div>
                  </div>   
            )
          })}
        </div>
        </div>
      )
    }    
