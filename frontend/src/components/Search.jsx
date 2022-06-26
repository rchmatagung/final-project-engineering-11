import React from 'react'

function Search() {
  return (
        <form className='bg-gray-200 p-5'>   
            <label className="mb-2 text-sm font-medium text-gray-900 sr-only">Search</label>
            <div className="relative mx-10 lg:mx-32">
                <div className="flex absolute inset-y-0 left-0 items-center pl-3 pointer-events-none">
                    <svg className="w-5 h-5 text-gray-500" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"></path></svg>
                </div>
                <input type="search" className="block p-4 pl-10 w-full text-sm text-gray-900 bg-gray-50 rounded-lg border border-gray-300 focus:ring-yellow-500 focus:border-yellow-500" placeholder="Search Mentor..." required />
                <button type="submit" className="text-white absolute right-2.5 bottom-2.5 bg-yellow-500 hover:bg-yellow-500 focus:ring-4 focus:outline-none focus:ring-yellow-300 font-medium rounded-lg text-sm px-4 py-2">Search</button>
            </div>
        </form>
  )
}

export default Search