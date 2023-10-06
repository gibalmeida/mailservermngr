<script setup lang="ts">
interface Props {
    id: string
    label: string
    placeholder?: string
    autocomplete?: string
    required?: boolean
    autofocus?: boolean
    disabled?: boolean
    class?: string
    modelValue: string
}
const props = defineProps<Props>()
const modelValue = props.modelValue
</script>
<template>
  <div class="mt-4">
    <label for="{{ props.id }}" class="font-semibold text-sm block">{{ props.label }}</label>
    <input
      type="text"
      :id="props.id"
      v-model=modelValue
      :placeholder="props.placeholder"
      :autocomplete="props.autocomplete"
      :required=props.required
      :autofocus="props.autofocus"
      :disabled="props.disabled"

      class="p-2 ring-1 rounded-sm w-full"
      :class="props.class"
    />
  </div>
</template>
