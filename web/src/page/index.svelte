<script>
    // TODO restrictions such as vegan/vegetarian, allergies, etc...
    // TODO snack/entree/dessert

    import Choices from "../../lib/Choices.svelte";
    import Svelecte from "svelecte";
    import {
        continents as continentMap,
        countries as countryMap,
        getEmojiFlag,
    } from "countries-list";

    const times = ["breakfast", "lunch", "dinner"];
    const flavors = ["sweet", "savory", "salty", "spicy", "sour", "bitter"];

    const continents = Object.values(continentMap)
        .sort()
        .map((c) => ({ id: c, name: c }))
        .filter((c) => c.id !== "Antarctica");

    const countries = Object.entries(countryMap)
        .sort(([, a], [, b]) => a.name.localeCompare(b.name))
        .map(([id, country]) => ({
            id: country.name,
            name: `${getEmojiFlag(id)}  ${country.name}`,
        }));

    let selectByContinent = false;

    let selectedTimes = [];
    let selectedFlavors = [];
    let selectedContinents = [];
    let selectedCountries = [];

    $: valid =
        selectedTimes.length > 0 &&
        selectedFlavors.length > 0 &&
        (selectByContinent ? selectedContinents.length > 0 : selectedCountries.length > 0);

    let includeIngredients;
    let excludeIngredients;

    let generating = false;

    async function generate() {
        generating = true;

        const resp = await fetch("/api/generate", {
            method: "POST",
            body: JSON.stringify({
                Times: selectedTimes,
                Flavors: selectedFlavors,
                InludeIngredients: includeIngredients,
                ExcludeIngredients: excludeIngredients,
                Locations: selectByContinent ? selectedContinents : selectedCountries,
            }),
        });


        if (resp.status === 200) {
            location.href = await resp.text();
        } else if (resp.status === 500) {
            alert("There was a problem generating recipe, try again.");
        } else {
            alert(`error ${resp.status}: ${await resp.text()}`);
        }

        generating = false;
    }
</script>

<svelte:head>
    <title>Plateful</title>
</svelte:head>

<div class="flex">
    <Choices choices={times} bind:selected={selectedTimes}>
        <h5>What kinds of meal are you craving?</h5>
    </Choices>
</div>

<div class="flex">
    <Choices choices={flavors} bind:selected={selectedFlavors}>
        <h5>What flavors do you want in your meal?</h5>
    </Choices>
</div>

<div class="flex">
    <h5>What ingredients do you want to use?</h5>
    <textarea
        bind:value={includeIngredients}
        placeholder="Type ingredients here"
    />
</div>

<div class="flex">
    <h5>What ingredients do you want to avoid?</h5>
    <textarea
        bind:value={excludeIngredients}
        placeholder="Type ingredients here"
    />
</div>

<div class="flex">
    <h5>Which types of cuisine do you want?</h5>
    {#if selectByContinent}
        <Svelecte
            placeholder="Select continents"
            options={continents}
            bind:value={selectedContinents}
            multiple
            style="max-width: 500px; width: 100%"
        ></Svelecte>
    {:else}
        <Svelecte
            placeholder="Select countries"
            options={countries}
            bind:value={selectedCountries}
            multiple
            style="max-width: 500px; width: 100%"
        ></Svelecte>
    {/if}
    <label>
        <input type="checkbox" role="switch" bind:checked={selectByContinent} />
        Select by continent instead
    </label>
</div>

<div class="flex">
    {#if generating}
        <button aria-busy="true">Recipe is being created...</button>
    {:else}
        <button on:click={generate} disabled={!valid}>Generate Recipe</button>
    {/if}
</div>

<style>
    .flex {
        padding: 20px 0;
        display: flex;
        align-items: center;
        flex-direction: column;
    }

    textarea {
        max-width: 500px;
        font-size: 16px;
        resize: none;
    }
</style>
