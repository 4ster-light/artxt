<script setup lang="ts">
useSeoMeta({
	title: "✰ArtSCII✰",
	description: "The best image to ascii art converter",
	ogTitle: "ArtSCII",
	ogDescription: "The best image to ascii art converter",
	ogImage: "/Icon.svg",
	ogUrl: "https://artscii.deno.dev",
	twitterTitle: "ArtSCII",
	twitterDescription: "The best image to ascii art converter",
	twitterImage: "/Icon.svg",
	twitterCard: "summary",
});

useHead({
	htmlAttrs: {
		lang: "en",
	},
	link: [
		{
			rel: "icon",
			type: "image/svg",
			href: "/Icon.svg",
		},
	],
});
</script>

<template>
  <main class="grow p-4 md:p-6 max-w-7xl mx-auto w-full">
    <div class="grid grid-cols-1 lg:grid-cols-12 gap-6">
      <div class="lg:col-span-4 space-y-6">
        <ImageUploader />
        <SettingsPanel />
      </div>
      <div class="lg:col-span-8">
        <AsciiOutput />
      </div>
    </div>
  </main>
</template>
