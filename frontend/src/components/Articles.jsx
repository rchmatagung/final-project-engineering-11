import React, { useState, useEffect } from 'react'
import Art from '../assets/support.jpg'
import api from '../api/api'

function Articles() {
  const [articles, setArticles] = useState([])
  const [showModal, setShowModal] = useState(false);

    const getArticle = async () => {
        try {
        await api.get('/user/artikel')
         .then((res) => {
          setArticles(res.data.data)
          console.log(res.data.data)
        })
        } catch (error) {
          console.log(error);
        }
    };

  useEffect(() => {
    getArticle()
  }, [])


  return (
    <div className="container mx-auto px-8 py-5">
      <h2 className='text-center text-4xl md:text-5xl hover:text-yellow-500 font-bold'>Artikel</h2>
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-10 space-x-3 mt-10">
        {articles.map((article) => {
          return (
            <div className='mt-4 col-span-1 bg-gray-200 shadow-xl' key={article.id}>
              <img className="rounded-t-lg" src={Art} alt="Article Pic"/>
              <div className="p-4">
                <div className='overflow-y-hidden h-52'>
                <h5 className="text-lg font-bold mb-2">{article.title}</h5>
                <p className="text-gray-600 text-sm leading-6">{article.content}</p>
                </div>
                <button
                  className="bg-yellow-500 hover:bg-yellow-700 text-white active:bg-yellow-600 font-bold hover:yellow-700 uppercase text-sm px-6 py-3 rounded shadow hover:shadow-lg outline-none focus:outline-none mr-1 mb-2 mt-5 ease-linear transition-all duration-150"
                  onClick={setShowModal(true)}>
                  Read More
                </button>
              {showModal  ? (
              <>
              <div
                className="justify-center items-center flex overflow-x-hidden overflow-y-auto fixed inset-0 z-50 outline-none focus:outline-none"
              >
                <div className="relative w-auto my-6 mx-auto max-w-3xl">
                  {/*content*/}
                  <div className="border-0 rounded-lg shadow-lg relative flex flex-col w-full bg-white outline-none focus:outline-none">
                    {/*header*/}
                    <div className="flex items-start justify-between p-5 border-b border-solid border-slate-200 rounded-t">
                      <h3 className="text-3xl font-semibold">
                        {article.title}
                      </h3>
                      <button
                        className="p-1 ml-auto bg-transparent border-0 text-black opacity-5 float-right text-3xl leading-none font-semibold outline-none focus:outline-none"
                        onClick={() => setShowModal(false)}
                      >
                        <span className="bg-transparent text-black opacity-5 h-6 w-6 text-2xl block outline-none focus:outline-none">
                          ×
                        </span>
                      </button>
                    </div>
                    {/*body*/}
                    <div className="relative p-6 flex-auto">
                      <p className="my-4 text-slate-500 text-lg leading-relaxed">
                      {article.content}
                      </p>
                    </div>
                    {/*footer*/}
                    <div className="flex items-center justify-end p-6 border-t border-solid border-slate-200 rounded-b">
                    <button
                    className="bg-yellow-500 text-white active:bg-emerald-600 font-bold uppercase text-sm px-6 py-3 rounded shadow hover:shadow-lg outline-none focus:outline-none mr-1 mb-1 ease-linear transition-all duration-150"
                    type="button"
                    onClick={() => setShowModal(false)}
                  >
                   Close
                  </button>
                    </div>
                  </div>
                </div>
              </div>
              <div className="opacity-25 fixed inset-0 z-40 bg-black"></div>
            </>
          ) : null}
                </div>
              </div>
            )
          })}
        </div>
    </div>
  )
}

export default Articles
