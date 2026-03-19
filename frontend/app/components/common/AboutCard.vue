<template>
<div class="grid grid-cols-4 gap-6 
    p-6 rounded-2xl shadow-xl
    bg-linear-to-br from-emerald-700/80 via-emerald-800/70 to-emerald-900/60">

    <div class="col-span-2 flex flex-col
        bg-white/10 backdrop-blur-md rounded-xl p-6 text-white min-h-70">

        <h2 class="text-xl font-semibold leading-snug">
            Система формирования заявок
        </h2>

        <p class="text-sm leading-relaxed text-emerald-100 mt-4">
            Информационная система предназначена для сбора, анализа и обработки заявок
            на подготовку кадров, а также формирования сводной потребности
            для учебных заведений.
        </p>
    </div>

    <div class="col-span-2 grid grid-cols-2 gap-5">
        <div class="bg-white/10 backdrop-blur-md rounded-xl p-5 flex flex-col min-h-40">
            <div class="flex items-center gap-3 mb-3">
                <Icon name="ph:student" class="w-7 h-7 text-white shrink-0" />
                <span class="text-base font-semibold text-white leading-tight">
                    Учебные заведения
                </span>
            </div>
            <p class="text-sm text-emerald-100 leading-relaxed wrap-break-word">
                Получают сводную информацию о потребности в специалистах и формируют планы подготовки кадров.
            </p>
        </div>

        <div class="bg-white/10 backdrop-blur-md rounded-xl p-5 flex flex-col min-h-40">
            <div class="flex items-center gap-3 mb-3">
                <Icon name="ph:buildings" class="w-7 h-7 text-white shrink-0" />
                <span class="text-base font-semibold text-white leading-tight">
                    Организации-заказчики
                </span>
            </div>
            <p class="text-sm text-emerald-100 leading-relaxed">
                Формируют заявки на подготовку специалистов и указывают требования к кадрам.
            </p>
        </div>

        <div class="bg-white/10 backdrop-blur-md rounded-xl p-5 flex flex-col min-h-40">
            <div class="flex items-center gap-3 mb-3">
                <Icon name="ph:user-gear" class="w-7 h-7 text-white shrink-0" />
                <span class="text-base font-semibold text-white leading-tight">
                    Администратор
                </span>
            </div>
            <p class="text-sm text-emerald-100 leading-relaxed">
                Обрабатывает заявки, выполняет анализ и формирует сводные отчёты.
            </p>
        </div>

        <div class="bg-white/10 backdrop-blur-md rounded-xl p-5 flex flex-col justify-between min-h-40">
            <div class="flex items-center gap-3 mb-3">
                <Icon name="ph:user" class="w-7 h-7 text-white shrink-0" />
                <span class="text-base font-semibold text-white leading-tight">
                    Специалисты
                </span>
            </div>
            <p class="text-sm text-emerald-100 leading-relaxed">
                Заполняют электронные формы и предоставляют данные для формирования заявок.
            </p>
        </div>

    </div>
</div>
</template>