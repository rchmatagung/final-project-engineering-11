import React, { useState, useEffect } from 'react'
import Art from '../assets/support.jpg'
import api from '../api/api'

function Articles() {
  const [articles, setArticles] = useState([])

    const getUser = async () => {
        try {
        await api.get('https://8ab4-111-94-105-152.ap.ngrok.io/user/artikel')
         .then((res) => {
          setArticles(res.data.data)
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
    <div className="container mx-auto px-8 py-5">
      <h2 className='text-center text-4xl md:text-5xl font-bold'>New Articles</h2>
      <div className="grid grid-cols-1 md:grid-cols-4 gap-10 pt-6 pb-5 px-20">
        {articles.map((article, index) => {
          return (
            <div className='md:col-span-2 lg:col-span-1' key={index}>
              <img className="rounded-t-lg" src={Art} alt="Article Pic"/>
              <div className="p-4">
                <h5 className="text-lg font-bold mb-2">{article.title}</h5>
                <p className="text-gray-600 text-sm leading-6">{article.content}</p>
                {/* <button type="button" className=" inline-block px-6 py-2.5 bg-blue-600 text-white font-medium text-xs leading-tight uppercase rounded shadow-md hover:bg-blue-700 hover:shadow-lg focus:bg-blue-700 focus:shadow-lg focus:outline-none focus:ring-0 active:bg-blue-800 active:shadow-lg transition duration-150 ease-in-out">Read More</button> */}
              </div>
            </div>
          )
        })}
      </div>
    </div>
  )
}

export default Articles