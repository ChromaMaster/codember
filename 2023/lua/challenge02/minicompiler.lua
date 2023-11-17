MiniCompiler = {
  Value = 0,
  Output = "",
}

function MiniCompiler:new(o)
  o = o or {}
  setmetatable(o, {
    __index = self,
  })
  return o
end

function MiniCompiler:tostring()
  return string.format("{Value: %d, Output: %s}", self.Value, self.Output)
end

function MiniCompiler:ParseInput(input)
  if input == '#' then
    self.Value = self.Value + 1
  elseif input == '@' then
    self.Value = self.Value - 1
  elseif input == '*' then
    self.Value = math.floor(self.Value ^ 2)
  elseif input == '&' then
    self.Output = string.format("%s%s", self.Output, tostring(self.Value))
  else
  end
end

return MiniCompiler
