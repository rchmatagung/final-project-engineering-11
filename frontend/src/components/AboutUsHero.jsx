import React from 'react'

const Hero = () => {
  return (
    <div className='w-full bg-zinc-200 flex flex-col justify-between'>
        <div className='bg-gray-100 max-w-[1024px] mx-8'>
            <div className='flex flex-col justify-center md:items-start w-full px-2 my-10'>
                <h1 className='m-auto py-3 text-4xl md:text-5xl lg:text-7xl font-bold'>About<span className="text-yellow-400">Us</span></h1>
                <p className='m-auto w-1/2 pb-5 text-center text-xs md:text-sm'>Hicoder akan membantu anda untuk mencari mentor yang berpengalaman di bidang IT. Anda akan dibimbing secara profesional dengan materi yang ter-update dan sesuai dengan kebutuhan perusahaan-perusahaan saat ini. Anda juga bisa menyampaikan pertanyaan dan melakukan konsultasi terkait kesulitan anda saat mempelajari materi atau mengerjakan project anda. Let's try with us.</p>
            </div>
        </div>
    </div>
  )
}

export default Hero