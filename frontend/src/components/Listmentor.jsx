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
        <div className=" bg-gray-200 dark:bg-white-600 flex flex-wrap items-center">
          {mentors.map((mentor, index) => {
                return (
          <div className="container max-w-lg bg-white rounded dark:bg-gray-800 shadow-lg transform duration-200 easy-in-out m-12"  key={index}>
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
                          src="https://ik.imagekit.io/t3defxmi6/unnamed_ESgTApzbR.png?ik-sdk-version=javascript-1.4.3&updatedAt=1654923915194"
                          alt="avatar"
                          className="mx-auto object-cover rounded-full h-24 w-24 bg-white p-1"
                        />
                      </span>
                    </div>
                    <div className="">
                      <div className="px-8 mb-8">
                        <h2 className="text-3xl font-bold text-green-800 dark:text-gray-300">
                        {mentor.name}
                        </h2>
                        <p className="text-gray-400 mt-2 dark:text-gray-400">
                        {mentor.skill}
                        </p>
                        <p className="mt-2 text-gray-600 dark:text-gray-300">
                          {mentor.bio}
                        </p>
                        <button onClick={() => window.location = ''} className=" mt-6 inline-block px-7 py-2 bg-yellow-500 cursor-pointer text-white font-medium text-xs leading-tight uppercase rounded shadow-md hover:bg-yellow-700 hover:shadow-lg focus:bg-yellow-700 focus:shadow-lg focus:outline-none focus:ring-0 active:bg-yellow-800 active:shadow-lg transition duration-150 ease-in-out">
                          Buat Janji
                        </button>
                      </div>
                    </div>
                  </div>
                  
            )
          })}
        </div>
      )
    }    
