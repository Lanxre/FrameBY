<template>
<div class="max-w-6xl mx-auto px-4 py-10 sm:py-14">
    <div 
      class="min-h-125 max-w-7xl mx-auto flex flex-col px-8 py-12 m-10 gap-8 
             shadow-2xl rounded-3xl overflow-hidden relative
             bg-cover bg-center text-white"
    >
      <div class="absolute inset-0 bg-linear-to-r from-emerald-600/70 via-emerald-700/60 to-emerald-800/40" />
  
      <div class="flex flex-col relative z-10 flex-1 gap-10">
        <h1 class="text-5xl flex items-center gap-2 md:text-6xl wrap-break-word w-190 font-bold uppercase leading-tight text-shadow-lg">
            Контакты
            <Icon name="ph:address-book" class="w-12 h-12 text-white" />
        </h1>
        <div class="flex flex-row items-center gap-2 text-2xl text-shadow-lg">
            <Icon name="ph:phone" class="w-12 h-12 text-white" />
            <span>+375 44 333 3333</span>
        </div>
        <div class="flex flex-row items-center gap-2 text-2xl text-shadow-lg">
            <Icon name="ph:telegram-logo" class="w-12 h-12 text-white" />
            <span>@Lanxre</span>
        </div>
      </div>
    </div>
</div>
</template>