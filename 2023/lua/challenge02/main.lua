local minicompiler = require("challenge02.minicompiler")

local inputFile = "challenge02/input"
local outputFile = "challenge02/output"

local function ReadFile(filename)
  local file = io.open(filename)
  if not file then return "" end

  local contents = file:read("a")

  file:close()

  return contents
end

local function WriteFile(filename, data)
  local file = io.open(filename, "w")
  if not file then error("could not open file") end
  file:write(data)
  file:close()
end


local fileContents = ReadFile(inputFile)

local compiler = minicompiler:new()

for i = 1, #fileContents do
  local c = fileContents:sub(i, i)
  compiler:ParseInput(c)
end

WriteFile(outputFile, compiler.Output)

print("Solution: " .. compiler.Output)
