local minicompiler = require("challenge02.minicompiler")

local inputFile = "challenge02/input"

local function ReadFile(filename)
  local file = io.open(filename)
  if not file then return "" end

  local contents = file:read("a")

  file:close()

  return contents
end

local fileContents = ReadFile(inputFile)

local compiler = minicompiler:new()

for i = 1, #fileContents do
  local c = fileContents:sub(i, i)
  compiler:ParseInput(c)
end

print("Solution: " .. compiler.Output)
