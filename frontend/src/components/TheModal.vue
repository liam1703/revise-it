<template>
  <div
    class="fixed w-screen h-screen z-100 inset-0 bg-grey-900/40 backdrop-blur"
  >
    <div
      class="fixed px-5 z-100 bg-white rounded-t-2xl sm:rounded-t-lg mx-auto my-auto translate-y-0 
      inset-0 items-center transition-all rounded-lg shadow-xl sm:px-12 flex flex-col justify-between overflow-y-hidden
      md:h-max  sm:max-h-max h-screen w-screen sm:max-w-3xl"
      aria-modal="true"
    >
      <div class="flex flex-col w-full py-7 overflow-hidden">
        <slot></slot>
      </div>
    </div>
</div>
</template>

<script setup>
</script>
