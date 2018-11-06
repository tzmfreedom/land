#!/usr/bin/env ruby

f = open('node.txt')
f.each do |line|
  klass = line.strip
#   puts <<~EOS
# func (n *#{klass}) Accept(v Visitor) interface{} {
#   return v.Visit#{klass}(n)
# }
#
# EOS
  puts <<~EOS
Visit#{klass}(Node) interface{}
EOS
end
