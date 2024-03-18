<template>
<nav class="flex justify-between p-4">
    <div class="italic font-medium text-sm">PAPER<span class="text-green-500 text-xs">LESS</span></div>
    <ul class="flex">
        <li class="mr-6">
        <RouterLink to="/" class=" hover:text-green-500">Search</RouterLink>
        </li>
        <li class="mr-6">
        <RouterLink to="/categories" class="hover:text-green-500">Categories</RouterLink>
        </li>
    </ul>
</nav>
</template>

<style>
.router-link-exact-active {
   color: #22c55e;
   font-weight: 700;
}
</style>