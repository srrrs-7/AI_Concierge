<script lang="ts">
const colors = [
  { name: "Red", hex: "#FF0000" },
  { name: "Green", hex: "#00FF00" },
  { name: "Blue", hex: "#0000FF" },
  { name: "Yellow", hex: "#FFFF00" },
  { name: "Purple", hex: "#800080" },
  { name: "Orange", hex: "#FFA500" },
  { name: "Pink", hex: "#FFC0CB" },
  { name: "Cyan", hex: "#00FFFF" },
  { name: "Magenta", hex: "#FF00FF" },
  { name: "Lime", hex: "#00FF00" },
  { name: "Teal", hex: "#008080" },
  { name: "Brown", hex: "#A52A2A" },
  { name: "Navy", hex: "#000080" },
  { name: "Silver", hex: "#C0C0C0" },
  { name: "Gray", hex: "#808080" },
  { name: "Maroon", hex: "#800000" },
  { name: "Olive", hex: "#808000" },
  { name: "Sky Blue", hex: "#87CEEB" },
  { name: "Indigo", hex: "#4B0082" },
  { name: "Turquoise", hex: "#40E0D0" },
  { name: "Gold", hex: "#FFD700" },
  { name: "Coral", hex: "#FF7F50" },
  { name: "Slate Gray", hex: "#708090" },
  { name: "Violet", hex: "#EE82EE" },
  { name: "Aqua", hex: "#00FFFF" },
  { name: "Chartreuse", hex: "#7FFF00" },
  { name: "Tomato", hex: "#FF6347" },
  { name: "Khaki", hex: "#F0E68C" },
  { name: "Sandy Brown", hex: "#F4A460" },
  { name: "Lavender", hex: "#E6E6FA" },
  { name: "Dark Green", hex: "#006400" },
  { name: "Salmon", hex: "#FA8072" },
  { name: "Thistle", hex: "#D8BFD8" },
  { name: "Deep Pink", hex: "#FF1493" },
  { name: "Dark Cyan", hex: "#008B8B" },
  { name: "Hot Pink", hex: "#FF69B4" },
  { name: "Lemon Chiffon", hex: "#FFFACD" },
  { name: "Peru", hex: "#CD853F" },
  { name: "Orchid", hex: "#DA70D6" },
  { name: "Powder Blue", hex: "#B0E0E6" },
  { name: "Dark Olive Green", hex: "#556B2F" },
  { name: "Light Coral", hex: "#F08080" },
  { name: "Medium Purple", hex: "#9370DB" },
  { name: "Crimson", hex: "#DC143C" },
  { name: "Dark Turquoise", hex: "#00CED1" },
  { name: "Deep Sky Blue", hex: "#00BFFF" },
  { name: "Light Goldenrod Yellow", hex: "#FAFAD2" },
  { name: "Chocolate", hex: "#D2691E" },
  { name: "Medium Orchid", hex: "#BA55D3" },
  { name: "Alice Blue", hex: "#F0F8FF" },
  { name: "Medium Sea Green", hex: "#3CB371" },
  { name: "Pale Goldenrod", hex: "#EEE8AA" },
  { name: "Sienna", hex: "#A0522D" },
  { name: "Medium Slate Blue", hex: "#7B68EE" },
  { name: "Light Sky Blue", hex: "#87CEFA" },
  { name: "Medium Aquamarine", hex: "#66CDAA" },
  { name: "Light Salmon", hex: "#FFA07A" },
  { name: "Lavender Blush", hex: "#FFF0F5" },
  { name: "Dark Orange", hex: "#FF8C00" },
  { name: "Dark Slate Blue", hex: "#483D8B" },
  { name: "Pale Green", hex: "#98FB98" },
  { name: "Medium Violet Red", hex: "#C71585" },
  { name: "Light Green", hex: "#90EE90" },
  { name: "Rosy Brown", hex: "#BC8F8F" },
  { name: "Deep Sky Blue", hex: "#00BFFF" },
  { name: "Pale Turquoise", hex: "#AFEEEE" },
  { name: "Light Pink", hex: "#FFB6C1" },
  { name: "Medium Spring Green", hex: "#00FA9A" },
  { name: "Light Steel Blue", hex: "#B0C4DE" },
  { name: "Indian Red", hex: "#CD5C5C" },
  { name: "Dark Orchid", hex: "#9932CC" },
  { name: "Antique White", hex: "#FAEBD7" },
  { name: "Spring Green", hex: "#00FF7F" },
  { name: "Honeydew", hex: "#F0FFF0" },
  { name: "Cornflower Blue", hex: "#6495ED" },
  { name: "Saddle Brown", hex: "#8B4513" },
  { name: "MediumTurquoise", hex: "#48D1CC" },
  { name: "Ghost White", hex: "#F8F8FF" },
  { name: "MediumBlue", hex: "#0000CD" },
  { name: "Ivory", hex: "#FFFFF0" },
  { name: "Dodger Blue", hex: "#1E90FF" },
  { name: "Dark Goldenrod", hex: "#B8860B" },
  { name: "Dark Slate Gray", hex: "#2F4F4F" },
  { name: "Light Yellow", hex: "#FFFFE0" },
  { name: "Steel Blue", hex: "#4682B4" },
  { name: "Peach Puff", hex: "#FFDAB9" },
  { name: "Light Blue", hex: "#ADD8E6" },
  { name: "MediumOrchid", hex: "#BA55D3" },
  { name: "Pale Violet Red", hex: "#DB7093" },
  { name: "Moccasin", hex: "#FFE4B5" },
  { name: "CornSilk", hex: "#FFF8DC" },
  { name: "Royal Blue", hex: "#4169E1" },
  { name: "Papaya Whip", hex: "#FFEFD5" },
  { name: "Medium Slate Blue", hex: "#7B68EE" },
  { name: "PaleGoldenrod", hex: "#EEE8AA" },
  { name: "Coral", hex: "#FF7F50" },
  { name: "Light Cyan", hex: "#E0FFFF" },
  { name: "MediumSpringGreen", hex: "#00FA9A" },
  { name: "Lawn Green", hex: "#7CFC00" },
  { name: "LightGrey", hex: "#D3D3D3" },
  { name: "Purple", hex: "#800080" },
  { name: "DeepPink", hex: "#FF1493" },
  { name: "Turquoise", hex: "#40E0D0" },
  { name: "DarkRed", hex: "#8B0000" },
  { name: "OliveDrab", hex: "#6B8E23" },
  { name: "Silver", hex: "#C0C0C0" },
  { name: "LightGray", hex: "#D3D3D3" },
  { name: "PaleGreen", hex: "#98FB98" },
  { name: "LightGreen", hex: "#90EE90" },
  { name: "LightCoral", hex: "#F08080" },
  { name: "Indigo", hex: "#4B0082" },
  { name: "Beige", hex: "#F5F5DC" },
  { name: "DarkKhaki", hex: "#BDB76B" },
  { name: "Aqua", hex: "#00FFFF" },
  { name: "DarkMagenta", hex: "#8B008B" },
  { name: "Silver", hex: "#C0C0C0" },
  { name: "DarkSalmon", hex: "#E9967A" },
  { name: "SlateBlue", hex: "#6A5ACD" },
  { name: "PaleTurquoise", hex: "#AFEEEE" },
  { name: "DarkViolet", hex: "#9400D3" },
  { name: "LightYellow", hex: "#FFFFE0" },
  { name: "LimeGreen", hex: "#32CD32" },
  { name: "Cornsilk", hex: "#FFF8DC" },
  { name: "LightSlateGray", hex: "#778899" },
  { name: "DarkGreen", hex: "#006400" },
  { name: "LightSkyBlue", hex: "#87CEFA" },
  { name: "MintCream", hex: "#F5FFFA" },
  { name: "DarkGray", hex: "#A9A9A9" },
  { name: "DarkCyan", hex: "#008B8B" },
  { name: "Black", hex: "#000000" },
];

function colorCheck(text: string, target: string): string {
  const regex = new RegExp(target, 'g');
  const matches = text.match(regex);
 if (matches ? matches.length : 0 > 0) {
    return "black";
  }
  return "white"
}

</script>

<h1>Color Palette</h1>
<div class="color__block">
  {#each colors as color}
    <div class="color" style="background-color: {color.hex}; color: {colorCheck(color.hex, "F")};">
      <p>{color.name}</p>
      <p>{color.hex}</p>
    </div>
  {/each}
</div>

<style>
  .color__block {
    background-color: rgb(206, 206, 206);
  }
  .color {
    display: inline-block;
    width: 180px;
    height: 80px;
    margin: 10px;
    padding: 10px;
    border-radius: 100px;
    border: 1px solid #000000;
    text-align: center;
    justify-content: center;
  }
</style>
