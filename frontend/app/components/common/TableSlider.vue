<script setup lang="ts">
import { ref } from 'vue'

const current = ref(0)

const slides = [
  'table1',
  'table2'
]

const next = () => {
  if (current.value < slides.length - 1) current.value++
}

const prev = () => {
  if (current.value > 0) current.value--
}
</script>

<template>
<div class="w-full space-y-4">

  <div class="flex items-center justify-between">
    <button
      @click="prev"
      class="p-2 rounded-xl bg-white/60 border border-black/5
             hover:bg-emerald-50 transition disabled:opacity-30"
      :disabled="current === 0"
    >
      <Icon name="mdi:chevron-left" class="w-5 h-5" />
    </button>

    <span class="text-sm text-gray-500">
      {{ current + 1 }} / {{ slides.length }}
    </span>

    <button
      @click="next"
      class="p-2 rounded-xl bg-white/60 border border-black/5
             hover:bg-emerald-50 transition disabled:opacity-30"
      :disabled="current === slides.length - 1"
    >
      <Icon name="mdi:chevron-right" class="w-5 h-5" />
    </button>
  </div>

  <!-- Slides -->
  <div class="relative overflow-hidden">

    <!-- Slide 1 -->
    <div v-if="current === 0"
            class="p-6 rounded-2xl bg-white/70 backdrop-blur border border-black/5 shadow-sm">

      <div class="text-center space-y-2 mb-6">
        <p class="font-semibold">
        СВЕДЕНИЯ  
        <br>
        о дополнительной потребности в молодых специалистах, рабочих, служащих
        <br>
        (с общим высшим, специальным высшим, средним специальным, профессионально-техническим образованием)
        </p>

        <div class="text-gray-400">______________________________</div>
        <p class="text-sm text-gray-500">(наименование организации)</p>

        <div class="text-gray-400">______________________________</div>
        <p class="text-sm text-gray-500">(место нахождения)</p>

        <div class="text-gray-400">______________________________</div>
        <p class="text-sm text-gray-500">(форма собственности)</p>
      </div>

      <div class="overflow-x-auto">
        <table class="w-full text-sm border border-black/10">
          <thead class="bg-emerald-50">
            <tr>
              <th class="border p-2 text-left">Специальность</th>
              <th class="border p-2 text-left">Квалификация</th>
              <th class="border p-2 text-left">Образование</th>
              <th class="border p-2 text-center">2025</th>
              <th class="border p-2 text-center">2026</th>
              <th class="border p-2 text-center">2027</th>
            </tr>
          </thead>
          <tbody>
            <tr>
              <td class="border p-2"></td>
              <td class="border p-2"></td>
              <td class="border p-2"></td>
              <td class="border p-2"></td>
              <td class="border p-2"></td>
              <td class="border p-2"></td>
            </tr>
            <tr class="bg-gray-50 font-medium">
              <td class="border p-2">Итого</td>
              <td class="border p-2"></td>
              <td class="border p-2"></td>
              <td class="border p-2"></td>
              <td class="border p-2"></td>
              <td class="border p-2"></td>
            </tr>
          </tbody>
        </table>
        <br>
        <div class="grid gap-4 text-sm">
          <div class="flex justify-between">
            <span>Руководитель организации</span>
            <span>______________________</span>
          </div>
  
          <div class="flex justify-between text-gray-500 text-xs">
            <span>(подпись)</span>
            <span>(инициалы, фамилия)</span>
          </div>
  
          <p class="text-xs text-gray-500 mt-4">
            * Для профессий рабочих с профессионально-техническим и средним специальным образованием.
          </p>
  
        </div>
      </div>
    </div>

    <!-- Slide 2 -->
    <div v-if="current === 1"
         class="p-6 rounded-2xl bg-white/70 backdrop-blur border border-black/5 shadow-sm">
             
             <div class="text-center space-y-2 mb-6">
               <p class="font-semibold">
               СВЕДЕНИЯ  
               <br>
               о дополнительной потребности в 
               <br>
                молодых специалистах с углубленным высшим образованием (степень – магистр)
               </p>
       
               <div class="text-gray-400">______________________________</div>
               <p class="text-sm text-gray-500">(наименование организации)</p>
       
               <div class="text-gray-400">______________________________</div>
               <p class="text-sm text-gray-500">(место нахождения организации)</p>
       
               <div class="text-gray-400">______________________________</div>
               <p class="text-sm text-gray-500">(наименование государственного органа, имеющего в подчинении учреждения образования, которые реализуют образовательную программу магистратуры)</p>
             </div>
       
             <div class="overflow-x-auto">
               <table class="w-full text-sm border border-black/10">
                 <thead class="bg-emerald-50">
                   <tr>
                     <th class="border p-2 text-left">Код и наименование специальности</th>
                     <th class="border p-2 text-left">Потребность по годам (выпуск из учреждений образования), человек</th>
                   </tr>
                 </thead>
                 <tbody>
                   <tr>
                     <td class="border p-2"></td>
                     <td class="border p-2"></td>
                   </tr>
                   <tr class="bg-gray-50 font-medium">
                     <td class="border p-2">Итого</td>
                     <td class="border p-2"></td>
                   </tr>
                 </tbody>
               </table>
               <br>
               <div class="grid gap-4 text-sm">
                 <div class="flex justify-between">
                   <span>Руководитель организации</span>
                   <span>______________________</span>
                 </div>
         
                 <div class="flex justify-between text-gray-500 text-xs">
                   <span>(подпись)</span>
                   <span>(инициалы, фамилия)</span>
                 </div>
               </div>
             </div>
    </div>

  </div>

</div>
</template>